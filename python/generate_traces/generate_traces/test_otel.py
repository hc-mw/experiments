from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.resources import Resource

from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter

# Set up the OTLP exporter
otlp_exporter = OTLPSpanExporter(
    endpoint="http://localhost:9320/v1/traces"
)

# otlp_exporter = OTLPSpanExporter(
#     endpoint="http://localhost:9320",  # Adjust if your collector is not on localhost
#     insecure=True  # Set to False if using TLS
# )

# Set up the trace provider
resource = Resource(attributes={
    "service.name": "test-service"
})
trace_provider = TracerProvider(resource=resource)
processor = BatchSpanProcessor(otlp_exporter)
trace_provider.add_span_processor(processor)
trace.set_tracer_provider(trace_provider)

# Get a tracer
tracer = trace.get_tracer(__name__)

# Create a span
with tracer.start_as_current_span("test-span"):
    print("Sending a test span")

# Flush the exporter
trace_provider.shutdown()